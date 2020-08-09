package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type Speciality struct {
	Index               int    `json:"index"`
	SpecialityGroupCode string `json:"speciality_group_code"` //专业组
	CollegeName         string `json:"college_name"`          //院校名称
	Province            string `json:"province"`              //省份
	City                string `json:"city"`                  //市级
	SpecialityCode      string `json:"speciality_code"`       //专业代码
	SpecialityName      string `json:"speciality_name"`       //专业名称
	SubjectEquirement   string `json:"subject_equirement"`    //科目要求
	LiberalArts         int    `json:"liberal_arts"`          //2019文
	LiberalArtsRank     int    `json:"liberal_arts_rank"`     //2019文排
	Science             int    `json:"science"`               //2019理
	ScienceRank         int    `json:"science_rank"`          //2019理排
	ComprehensiveRank   int    `json:"comprehensive_rank"`    //2020综合位次
	Engineering985      int    `json:"engineering_985"`       //985工程 0：非 1：是
	Engineering211      int    `json:"engineering_211"`       //211工程 0：非 1：是
	FirstClassCollege   int    `json:"first_class_college"`   //双一流高校 0：非 1：是
	FirstClassSubject   int    `json:"first_class_subject"`   //双一流学科 0：非 1：是
	Attribute           int    `json:"attribute"`             //属性 0：省属 1：部属
	Department          string `json:"department"`            //直属部门
	Type                string `json:"type"`                  //院校类型
	Memo                string `json:"memo"`                  //专业备注
	PlanNumber          int    `json:"plan_number"`           //计划数
	LengthOfSchooling   string `json:"length_of_schooling"`   //学制
	StudyExpenses       int    `json:"study_expenses"`        //收费标准
	Year                int    `json:"year"`                  //年份
	Grade               int    `json:"grade"`                 //等级 0:A 1:B 2:A提前 3:B提前
	New                 int    `json:"new"`                   //新增 0：非 1：是
}

//データベースの初期化　Init
func Init() error {
	// DSN:Data Source Name
	var err error
	dsn := "root:root123456@tcp(127.0.0.1:3307)/college_info"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}

//大学名で専攻を取得
func GetSpecialityByCollegeName(name string) ([]Speciality, error) {
	rows, err := db.Query("SELECT * FROM base_info where college_name=?", name)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	specialitys := make([]Speciality, 0)

	for rows.Next() {
		var speciality Speciality
		rows.Scan(&speciality.Index,
			&speciality.SpecialityGroupCode,
			&speciality.CollegeName,
			&speciality.Province,
			&speciality.City,
			&speciality.SpecialityCode,
			&speciality.SpecialityName,
			&speciality.SubjectEquirement,
			&speciality.LiberalArts,
			&speciality.LiberalArtsRank,
			&speciality.Science,
			&speciality.ScienceRank,
			&speciality.ComprehensiveRank,
			&speciality.Engineering985,
			&speciality.Engineering211,
			&speciality.FirstClassCollege,
			&speciality.FirstClassSubject,
			&speciality.Attribute,
			&speciality.Department,
			&speciality.Type,
			&speciality.Memo,
			&speciality.PlanNumber,
			&speciality.LengthOfSchooling,
			&speciality.StudyExpenses,
			&speciality.Year,
			&speciality.Grade,
			&speciality.New)
		specialitys = append(specialitys, speciality)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return specialitys, nil
}
