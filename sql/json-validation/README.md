# JSON Validation

## example

```sql
-- 建表,设置校验
CREATE TABLE geo (
    coordinate JSON,
  CHECK(
        JSON_SCHEMA_VALID(
            '{
                "type":"object",
                "properties":{
                      "latitude":{"type":"number", "minimum":-90, "maximum":90},
                      "longitude":{"type":"number", "minimum":-180, "maximum":180}
                },
                "required": ["latitude", "longitude"]
            }',
            coordinate
        )
   )
)ENGINE = InnoDB AUTO_INCREMENT = 10000 DEFAULT CHARSET = utf8mb4 COMMENT = '地理信息';

INSERT INTO `geo` (`coordinate`) VALUES ('{"latitude":80,"longitude":90}');
```

修改已有表的校验

```sql
-- car_info表的car_detail字段增加校验
ALTER TABLE car_info ADD CONSTRAINT CHECK(
  JSON_SCHEMA_VALID(
    '{
      "$schema": "http://json-schema.org/draft-04/schema#",
      "type": "object",
      "properties": {
        "tags": {
          "type": "array",
          "items": { "type": "string" }
        }
      },
      "additionalProperties": false
    }',
    car_detail
  )
);
```

## references

- [sql for devs](https://sqlfordevs.com/json-schema-validation)
- [mysql doc](https://dev.mysql.com/doc/refman/8.0/en/json-validation-functions.html#function_json-schema-valid)
