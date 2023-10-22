package lnt

var GetCurrentLnt = `Select id_course,type_course,name_of_course,description_of_course
	from hackerevolution.users
    inner join hackerevolution.user_courses on users.id_user = user_courses.id_user
    inner join hackerevolution.courses c on c.id_course = user_courses.id_c where users.id_user = $1`

var GetCurrentNode = `Select note from hackerevolution.user_note_course
inner join hackerevolution.users u on u.id_user = user_note_course.id_user
inner join hackerevolution.courses c on c.id_course = user_note_course.id_course
inner join hackerevolution.user_courses uc on c.id_course = uc.id_c
where c.id_course = $2 and u.id_user=$1`

var CheckExistUserToCourse = `Select 
    Exists( Select * from hackerevolution.user_courses where id_user = $1 and id_c = $2)`

var GetAllZapis = `Select id_post,id_task,text,created_at,updated_at from hackerevolution.posts where id_course = $1`

var SendOtchet = `Select fio, mark,max_mark, text from hackerevolution.users u
join hackerevolution.task_user tu
on id_user = tu.id_u join hackerevolution.tasks t on tu.id_t = t.id_task
join hackerevolution.user_courses uc on u.id_user = uc.id_user
join hackerevolution.courses c on uc.id_c = c.id_course
where id_course= $1`
