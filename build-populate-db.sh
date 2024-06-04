FILE='./lib-babel.db'

if [ -f $FILE ]
then
		echo 'the file already exists'
		echo 'do you want to redo?(Y/n) :'
		read option
		case $option in
				"n"|"N")
						exit 0;
						;;
				*)
						rm $FILE
						;;
		esac
fi

# db file does not exist 
cat ./sqlite3-scripts/create-tables.sql | sqlite3 $FILE	
cat ./sqlite3-scripts/users.sql | sqlite3 $FILE	
cat ./sqlite3-scripts/books.sql | sqlite3 $FILE	
echo "completed"
