# initialize database
mysql -u root -p < ./docs/db/melisprint_db.sql
# create user
mysql -u root -p < ./docs/db/melisprint_db_user.sql