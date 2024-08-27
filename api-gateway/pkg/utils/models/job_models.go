package models

import (
	"time"
)

type JobOpening struct {
	ID                  uint      `json:"id"`
	Title               string    `json:"title"`
	Description         string    `json:"description"`
	Requirements        string    `json:"requirements"`
	Location            string    `json:"location"`
	EmploymentType      string    `json:"employment_type"`
	Salary              string    `json:"salary"`
	SkillsRequired      string    `json:"skills_required"`
	ExperienceLevel     string    `json:"experience_level"`
	EducationLevel      string    `json:"education_level"`
	ApplicationDeadline time.Time `json:"application_deadline"`
}

type JobOpeningResponse struct {
	ID                  uint      `json:"id"`
	Title               string    `json:"title"`
	Description         string    `json:"description"`
	Requirements        string    `json:"requirements"`
	PostedOn            time.Time `json:"posted_on"`
	EmployerID          int32     `json:"employer_id"`
	Location            string    `json:"location"`
	EmploymentType      string    `json:"employment_type"`
	Salary              string    `json:"salary"`
	SkillsRequired      string    `json:"skills_required"`
	ExperienceLevel     string    `json:"experience_level"`
	EducationLevel      string    `json:"education_level"`
	ApplicationDeadline time.Time `json:"application_deadline"`
}

type AllJob struct {
	ID                  uint      `json:"id"`
	Title               string    `json:"title"`
	ApplicationDeadline time.Time `json:"application_deadline"`
	EmployerID          int32     `json:"employer_id"`
}

type JobSeekerGetAllJobs struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
}

type ApplyJob struct {
	JobseekerID int64  `json:"jobseeker_id" validate:"required"`
	JobID       int64  `json:"job_id" validate:"required"`
	Resume      []byte `json:"resume" validate:"required"`
	ResumeURL   string `json:"resume_url" validate:"required"`
	CoverLetter string `json:"cover_letter" validate:"lte=500"`
}

type ApplyJobResponse struct {
	ID          uint   `json:"id"`
	JobseekerID int64  `json:"jobseeker_id" validate:"required"`
	JobID       int64  `json:"job_id" validate:"required"`
	ResumeURL   string `json:"resume_url" validate:"required"`
	CoverLetter string `json:"cover_letter" validate:"lte=500"`
}