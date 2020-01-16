for i in 'find -name "mys*" | wc -l'
do
	cd the-final-cl-test/mystery/
	export SUS='head -n 179 streets/Buckingham_Place | tail -n 1 | cut -d '#' -f2'
	echo $SUS
	head interviews/interview-$SUS
	grep -B 3 -A 1 "Erika Owens\|Joe Germusha\|Dartey Henv\|Hellen Maher" vehicles
	grep -l $MAIN_SUSPECT * | wc -l
done
