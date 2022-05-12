package associate

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/benny502/gorm-helper/builder"
	"gorm.io/gorm/schema"
)

type associate struct {
	preload string
	joins   []string
}

func (assoc *associate) GetPreload() string {
	return assoc.preload
}

func (assoc *associate) GetJoinsString() []string {
	return assoc.joins
}

func NewAssociate(model schema.Tabler, preload string) builder.Associate {
	preloads := strings.Split(preload, ".")
	joins = make([]*joinStatement, 0)
	joinStat := slice(model, preloads[0], preloads[1:]...)
	joinStr := make([]string, 0)
	for _, stat := range joinStat {
		sql := fmt.Sprintf("LEFT JOIN %s ON %s.%s=%s.%s", stat.ForeignTable, stat.TableName, stat.PrimaryField, stat.ForeignTable, stat.ForeignField)
		joinStr = append(joinStr, sql)
	}
	return &associate{
		preload: preload,
		joins:   joinStr,
	}
}

type joinStatement struct {
	Statement    string
	ClassName    string
	TableName    string
	Primary      string
	PrimaryField string
	ForeignClass string
	ForeignTable string
	Foreign      string
	ForeignField string
}

var joins []*joinStatement

func slice(src schema.Tabler, preload string, others ...string) []*joinStatement {
	stat := &joinStatement{
		Primary:      "Id",
		PrimaryField: "id",
	}

	obj, ok := getStructRelateObject(src, preload, stat)
	if ok {
		joins = append(joins, stat)
	}

	if len(others) > 0 {
		slice(obj, others[0], others[1:]...)
	}
	return joins
}

func getStructRelateObject(src schema.Tabler, target string, stat *joinStatement) (schema.Tabler, bool) {
	if src == nil {
		return nil, false
	}

	stat.TableName = src.TableName()

	stat.ForeignClass = target

	reflectType := reflect.TypeOf(src)
	if reflectType.Kind() == reflect.Ptr {
		reflectType = reflectType.Elem()
	}

	if reflectType.Kind() == reflect.Struct {
		stat.ClassName = reflectType.Name()
		if result, ok := reflectType.FieldByName(target); ok {
			if primary, ok := getGormTag(result, "foreignKey"); ok {
				stat.Primary = primary
				if primaryField, ok := reflectType.FieldByName(primary); ok {
					if primaryName, ok := getGormTag(primaryField, "column"); ok {
						stat.PrimaryField = primaryName
					}
				}
			}

			if reference, ok := getGormTag(result, "references"); ok {
				stat.Foreign = reference
			} else {
				//默认用关联表的类名+Id做引用
				stat.Foreign = fmt.Sprintf("%s%s", stat.ClassName, stat.Primary)
			}

			if foreignField, ok := result.Type.FieldByName(stat.Foreign); ok {
				if foreignName, ok := getGormTag(foreignField, "column"); ok {
					stat.ForeignField = foreignName
				}
			}

			if obj, ok := reflect.New(result.Type).Interface().(schema.Tabler); ok {
				stat.ForeignTable = obj.TableName()
				return obj, true
			}

		}
	}

	return nil, false
}

func getGormTag(field reflect.StructField, lookup string) (string, bool) {
	if gorm, ok := field.Tag.Lookup("gorm"); ok {
		tags := strings.Split(gorm, ";")
		for _, tag := range tags {
			keyPair := strings.Split(tag, ":")
			if len(keyPair) == 2 && keyPair[0] == lookup {
				return keyPair[1], true
			}
		}
	}
	return "", false
}
