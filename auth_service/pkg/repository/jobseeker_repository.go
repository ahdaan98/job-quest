package repository

import (
	"Auth/pkg/domain"
	interfaces "Auth/pkg/repository/interface"
	"Auth/pkg/utils/models"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type jobSeekerRepository struct {
	DB *gorm.DB
}

func NewJobSeekerRepository(DB *gorm.DB) interfaces.JobSeekerRepository {
	return &jobSeekerRepository{
		DB: DB,
	}
}

func (jr *jobSeekerRepository) JobSeekerSignUp(jobSeekerDetails models.JobSeekerSignUp) (models.JobSeekerDetailsResponse, error) {
	var model models.JobSeekerDetailsResponse

	fmt.Println("email", model.Email)

	fmt.Println("models", model)
	if err := jr.DB.Raw("INSERT INTO job_seekers (email, password, first_name, last_name, phone_number, date_of_birth, gender) VALUES (?, ?, ?, ?, ?, ?, ?) RETURNING id, email, first_name, last_name, phone_number, date_of_birth, gender", jobSeekerDetails.Email, jobSeekerDetails.Password, jobSeekerDetails.FirstName, jobSeekerDetails.LastName, jobSeekerDetails.PhoneNumber, jobSeekerDetails.DateOfBirth, jobSeekerDetails.Gender).Scan(&model).Error; err != nil {
		return models.JobSeekerDetailsResponse{}, err
	}
	fmt.Println("inside", model.Email)
	return model, nil
}

func (jr *jobSeekerRepository) CheckJobSeekerExistsByEmail(email string) (*domain.JobSeeker, error) {
	var jobSeeker domain.JobSeeker
	res := jr.DB.Where(&domain.JobSeeker{Email: email}).First(&jobSeeker)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return &domain.JobSeeker{}, res.Error
	}
	return &jobSeeker, nil
}

func (jr *jobSeekerRepository) FindJobSeekerByEmail(jobSeeker models.JobSeekerLogin) (models.JobSeekerSignUp, error) {
	var user models.JobSeekerSignUp
	err := jr.DB.Raw("SELECT * FROM job_seekers WHERE email=? ", jobSeeker.Email).Scan(&user).Error
	if err != nil {
		return models.JobSeekerSignUp{}, errors.New("error checking user details")
	}
	return user, nil
}

func (jr *jobSeekerRepository) SaveLinkedinCredentials(jobSeekerDetails models.JobSeekerDetailsResponse) (models.JobSeekerDetailsResponse, error) {
    var model models.JobSeekerDetailsResponse

    // Debugging prints to check values
    fmt.Println("email:", jobSeekerDetails.Email)
    fmt.Println("models:", jobSeekerDetails)

    // Using raw SQL to insert and return the newly inserted row
    query := `INSERT INTO job_seekers (email, first_name, last_name, phone_number, date_of_birth, gender) 
              VALUES (?, ?, ?, ?, ?, ?) 
              RETURNING id, email, first_name, last_name, phone_number, date_of_birth, gender`

    // Execute the query and scan the result into the model
    if err := jr.DB.Raw(query, jobSeekerDetails.Email, jobSeekerDetails.FirstName, jobSeekerDetails.LastName, jobSeekerDetails.PhoneNumber, jobSeekerDetails.DateOfBirth, jobSeekerDetails.Gender).Scan(&model).Error; err != nil {
        return models.JobSeekerDetailsResponse{}, err
    }

    // Debugging print to check the result
    fmt.Println("inside:", model.Email)

    return model, nil
}

func (jr *jobSeekerRepository) IsJobSeekerExist(email string) bool {
    var count int64
    // Count the number of records matching the email
	query := `SELECT COUNT(*) FROM job_seekers WHERE email = ?`
    if err := jr.DB.Raw(query, email).Scan(&count).Error; err != nil {
        fmt.Println("error checking if job seeker exists:", err)
        return false
    }
    // Return true if count is greater than 0, meaning the job seeker exists
    return count > 0
}

func (jr *jobSeekerRepository) GetJobSeeker(email string) (models.JobSeekerDetailsResponse, error) {
    var jobSeeker models.JobSeekerDetailsResponse

    // Define the raw SQL query
    query := `SELECT id, email, first_name, last_name, phone_number, date_of_birth, gender FROM job_seekers WHERE email = ?`
    
    // Execute the query and scan the result into jobSeeker
    err := jr.DB.Raw(query, email).Scan(&jobSeeker).Error
    if err != nil {
        // Check if the error is due to record not being found
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return models.JobSeekerDetailsResponse{}, nil
        }
        // Return any other error
        return models.JobSeekerDetailsResponse{}, err
    }

    // Print the retrieved job seeker details for debugging
    fmt.Printf("Retrieved Job Seeker: %+v\n", jobSeeker)

    // Check if fields are properly populated
    if jobSeeker.Email == "" {
        fmt.Println("Warning: Email field is empty")
    }
    if jobSeeker.FirstName == "" {
        fmt.Println("Warning: FirstName field is empty")
    }
    if jobSeeker.LastName == "" {
        fmt.Println("Warning: LastName field is empty")
    }
    if jobSeeker.PhoneNumber == "" {
        fmt.Println("Warning: PhoneNumber field is empty")
    }
    if jobSeeker.DateOfBirth == "" {
        fmt.Println("Warning: DateOfBirth field is empty")
    }
    if jobSeeker.Gender == "" {
        fmt.Println("Warning: Gender field is empty")
    }

    return jobSeeker, nil
}

func (jr *jobSeekerRepository) GetJobSeekerEmailByID(id uint) (string, error) {
    var email string
    query := "SELECT email FROM job_seekers WHERE id = ?"
    if err := jr.DB.Raw(query, id).Scan(&email).Error; err != nil {
        return "", fmt.Errorf("error retrieving email: %v", err)
    }

    if email == "" {
        return "", fmt.Errorf("no email found for job seeker ID: %d", id)
    }

    return email, nil
}

func (jr *jobSeekerRepository) ActivateJobSeekerSubscriptionByPlanID(jobSeekerID uint, subscriptionPlanID uint) error {
    var plan models.JobseekerPlan
    var existingSubscription models.JobseekerSubscription

    // 1. Fetch the subscription plan details
    planQuery := "SELECT id, apply_jobs FROM jobseeker_plans WHERE id = ?"
    if err := jr.DB.Raw(planQuery, subscriptionPlanID).Scan(&plan).Error; err != nil {
        return fmt.Errorf("subscription plan not found: %v", err)
    }

    // 2. Check if there is already an active subscription for this job seeker
    subscriptionQuery := `
        SELECT id FROM jobseeker_subscriptions 
        WHERE job_seeker_id = ? AND is_active = true AND end_date > ?
        LIMIT 1`
    if err := jr.DB.Raw(subscriptionQuery, jobSeekerID, time.Now()).Scan(&existingSubscription).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
        return fmt.Errorf("error checking existing subscription: %v", err)
    }

    if existingSubscription.ID != 0 {
        return fmt.Errorf("subscription plan is already active for job seeker ID %d", jobSeekerID)
    }

    // 3. Calculate start date and end date (1 month from now)
    startDate := time.Now()
    endDate := startDate.AddDate(0, 1, 0) // Adds 1 month to the start date

    // 4. Insert a new subscription record
    insertQuery := `
        INSERT INTO jobseeker_subscriptions (job_seeker_id, subscription_plan_id, remaining_job_applications, start_date, end_date, is_active)
        VALUES (?, ?, ?, ?, ?, ?)`
    if err := jr.DB.Exec(insertQuery, jobSeekerID, subscriptionPlanID, plan.ApplyJobs, startDate, endDate, true).Error; err != nil {
        return fmt.Errorf("failed to activate job seeker subscription: %v", err)
    }

    return nil
}

func (jr *jobSeekerRepository) IsJobSeekerPlanActive(jobSeekerID uint) (bool, error) {
    var subscription domain.JobseekerSubscription

    err := jr.DB.Where("job_seeker_id = ? AND is_active = ? AND end_date > ?", jobSeekerID, true, time.Now()).
        First(&subscription).Error

    if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return false, nil
        }
        return false, fmt.Errorf("error checking if job seeker plan is active: %v", err)
    }

    return true, nil
}