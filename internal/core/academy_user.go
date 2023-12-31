package core

import (
	"fmt"
	db "github.com/VATUSA/api-v3/internal/database"
	"github.com/VATUSA/api-v3/pkg/constants"
)

func CreateAcademyUser(au *db.AcademyUser) error {
	// TODO
	return nil
}

func SyncAcademyUser(au *db.AcademyUser) error {
	// TODO
	return nil
}

func SyncRoles(a *db.AcademyUser) error {
	err := clearRoles(a)
	if err != nil {
		return err
	}

	err = addFacilityRole(a, constants.RoleStudent, a.Controller.Facility)
	if err != nil {
		return err
	}

	if IsSeniorStaff(a.Controller, a.Controller.Facility) {
		err = addFacilityRole(a, constants.RoleInstructor, "USA")
		if err != nil {
			return err
		}
		err = addFacilityRole(a, constants.RoleFacilityAdmin, a.Controller.Facility)
		if err != nil {
			return err
		}
	}

	if HasRole(a.Controller, constants.Instructor, a.Controller.Facility) {
		err = addFacilityRole(a, constants.RoleInstructor, "USA")
		if err != nil {
			return err
		}
		err = addFacilityRole(a, constants.RoleInstructor, a.Controller.Facility)
		if err != nil {
			return err
		}
	}

	for _, v := range a.Controller.Visits {
		err = addFacilityRole(a, constants.RoleStudent, v.Facility)
		if err != nil {
			return err
		}

		if HasRole(a.Controller, constants.Instructor, v.Facility) {
			err = addFacilityRole(a, constants.RoleInstructor, v.Facility)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func clearRoles(a *db.AcademyUser) error {
	// TODO
	return nil
}

func addFacilityRole(a *db.AcademyUser, roleID uint64, facility constants.Facility) error {
	return addCategoryRole(a, roleID, facility)
}

func addCategoryRole(a *db.AcademyUser, roleID uint64, category string) error {
	// TODO
	return nil
}

func addCourseRole(a *db.AcademyUser, roleID uint64, courseID uint64) error {
	// TODO
	return nil
}

func SyncCohorts(a *db.AcademyUser) error {
	var cohorts []string

	if !a.Controller.IsActive || a.Controller.Facility == constants.InactiveFacility {
		return nil
	}

	// Home Facility Cohort
	cohorts = append(cohorts, a.Controller.Facility)

	if HasRole(a.Controller, "MTR", a.Controller.Facility) {
		cohorts = append(cohorts, "TNG")
		cohorts = append(cohorts, "MTR")
		cohorts = append(cohorts, fmt.Sprintf("%s-MTR", a.Controller.Facility))
	}
	if HasRole(a.Controller, "INS", a.Controller.Facility) {
		cohorts = append(cohorts, "TNG")
		cohorts = append(cohorts, "INS")
		cohorts = append(cohorts, fmt.Sprintf("%s-INS", a.Controller.Facility))
	}

	// Visitor Cohorts
	for _, v := range a.Controller.Visits {
		cohorts = append(cohorts, fmt.Sprintf("%s-V", v.Facility))
		if HasRole(a.Controller, "MTR", v.Facility) {
			cohorts = append(cohorts, fmt.Sprintf("%s-MTR", v.Facility))
		}
		if HasRole(a.Controller, "INS", v.Facility) {
			cohorts = append(cohorts, fmt.Sprintf("%s-INS", v.Facility))
		}
	}

	// Rating Cohort
	if constants.Rating(a.Controller.ATCRating) > constants.C1 {
		cohorts = append(cohorts, "C1+")
		cohorts = append(cohorts, fmt.Sprintf("%s-C1+", a.Controller.Facility))
	} else {
		cohorts = append(cohorts, constants.RatingShort(a.Controller.ATCRating))
		cohorts = append(cohorts, fmt.Sprintf(
			"%s-%s", a.Controller.Facility, constants.RatingShort(a.Controller.ATCRating)))
	}

	return setUserCohorts(a, cohorts)
}

func setUserCohorts(a *db.AcademyUser, cohorts []string) error {
	err := clearCohorts(a)
	if err != nil {
		return err
	}

	for _, c := range cohorts {
		err = addCohort(a, c)
		if err != nil {
			return err
		}
	}
	return nil
}

func clearCohorts(a *db.AcademyUser) error {
	// TODO
	return nil
}

func addCohort(a *db.AcademyUser, cohort string) error {
	// TODO
	return nil
}
