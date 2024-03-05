# DIRECTORY HINT 



# Available SQL Variables (and their golang alias)
realisation in ./converter/

Numeric : TINYINT  (uint8)
          SMALLINT (int16)
          INT      (int32)
          BIGINT   (int64)
            

Aproximate: FLOAT (float32)
            REAL  (float64)
            DECIMAL 


Characters: CHAR           (string with limit)
            VARCHAR        (string with limit)
            VARCHAR(limit) (string with limit)
            TEXT           (string)


Time: DATE     (Date struct)
      TIME     (Time struct)
      DATETIME (DateTime struct)


# Available SQL Functions
realisation in ./converter/functions.go




