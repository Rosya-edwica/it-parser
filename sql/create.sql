CREATE TABLE platform(
    id SERIAL,
    name VARCHAR(255) NOT NULL,
    logo TEXT NOT NULL,
    url VARCHAR(255) NOT NULL,

    primary key(id)
);

CREATE TABLE vacancy(
    id varchar(255) not null,
    id_platform int NOT NULL,
    name varchar(255) not null,
    description text not null,
    url text not null,
    skills text[],
    salary_to int,
    salary_from int,
    currency varchar(3) DEFAULT 'RUB',
    city varchar(255) not null,
    address text not null,
    experience varchar(20) check(experience in (
        'Intern',
        'Junior',
        'Middle',
        'Senior',
        'Lead'
    )),
    schedule varchar(20) check(schedule in (
        'Полный день',
        'Сменный график',
        'Гибкий график',
        'Удаленная работа',
        'Вахтовый метод'
    )),
    employment varchar(20) check(employment in (
        'Полная занятость',
        'Частичная занятость',
        'Проектная работа/разовое задание',
        'Волонтерство',
        'Стажировка'
    )),
    primary key(id, id_platform),
    foreign key(id_platform) references platform(id) on delete cascade
);


