First of all make sure you have local mysql database connection running on port 3306,
database with the name "mydb" and a "promotion" table in that database (if the port
number or database name differ you can change it easily from repository/database.go
file in "dataSourceName" variable.
Also you'll need to write your credentials of mysql database in the same file appropriately
to "userName" and "password".

Run in you mysql command shell the following command: SET GLOBAL local_infile=1;

When you run the application all the entries from .csv file will be imported to the database.
You will have an endpoint, which will return a promotion by its "id" field.

Example: http://localhost:8081/promotion/0006c161-b9d2-4b62-988c-c25255a20965

Expected result: {
                     "id": "0006c161-b9d2-4b62-988c-c25255a20965",
                     "price": 31.457293,
                     "expiration_date": "2018-06-24T12:50:03Z"
                 }
