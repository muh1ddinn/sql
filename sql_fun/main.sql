
=========================================
CREATE OR REPLACE FUNCTION most_user_year()
RETURNS integer
AS $$
BEGIN 
RETURN (SELECT EXTRACT(YEAR FROM created_ts) AS year_count
FROM users
GROUP BY year_count
ORDER BY COUNT(*)DESC
);
END;
$$ LANGUAGE plpgsql;


=============================================
CREATE OR REPLACE FUNCTION get_users_by_birth_month(input_month int)
RETURNS TABLE (
    user_id int,
    user_first_name varchar,
    birth_date DATE
) AS
$$
BEGIN
    RETURN QUERY
    SELECT
        users.user_id,
        users.user_first_name,
        users.user_dob
    FROM
        users
    WHERE
        EXTRACT(MONTH FROM users.user_dob) IN (6, 7, 8);
END;
$$ LANGUAGE plpgsql;


===================================================
CREATE OR REPLACE FUNCTION userr_in_years(input_year int)
RETURNS TABLE (
    user_first_name varchar,
    user_dob DATE
) AS
$$
BEGIN
    RETURN QUERY
    SELECT
        users.user_first_name,
        users.user_dob
    FROM
        users
    WHERE
        EXTRACT(YEAR FROM users.user_dob) = input_year;
END;
$$ LANGUAGE plpgsql;=

========================================================== 
CREATE OR REPLACE FUNCTION genderuser(gender varchar)       ||
RETURN INTEGER  ||
LANGUAGE plpgsql  ||
AS  ||
$$  ||
BEGIN  ||
  RETURN(  ||
  SELECT count(user_gender)  ||
  FROM users WHERE user_gender=gender);  ||
END;  ||
$$;  ||