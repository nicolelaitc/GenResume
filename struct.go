package main

type Resume struct {
	Header Header
	// each of the section below is optional
	Summary               string
	Education             []Education
	WorkExperiences       []WorkExperience
	AdditionalExperiences []WorkExperience
	Skills                []SkillCategory
	Projects              []Project
	Certifications        []Certification
}

type Header struct {
	Name     string
	City     string
	Province string
	Email    string
	Linkedin string
	Github   string
	Phone    string
	Website  string
}

type Education struct {
	Institution string
	Degree      string
	Start       string
	End         string
	City        string
	Province    string
	Description []string
}

type WorkExperience struct {
	Company     string
	Position    string
	Start       string
	End         string
	City        string
	Province    string
	Description []string
}

type SkillCategory struct {
	Category string
	Skill    []string
}

type Project struct {
	Name         string
	Description  []string
	Link         string
	Technologies []string
}

type Certification struct {
	Name        string
	Institution string
	Date        string
}
