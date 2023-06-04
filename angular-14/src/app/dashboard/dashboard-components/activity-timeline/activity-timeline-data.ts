export interface Activity{
    name:string;
    image:string;
    rating:number;
    commentTime:string;
    comment:string;
}

export const activities:Activity[]=[
    {
        name:'Abay Alen',
        image:'assets/images/users/1.jpg',
        commentTime:'2023-05-28',
        rating:5,
        comment:'Взял курс для начинающих по бекенд разработке на Golang, доволен всем! Курсы очень обучающие',

    },
    {
        name:'Ainur Kozha',
        image:'assets/images/users/2.jpg',
        commentTime:'2023-05-27',
        rating:4,
        comment:'"Прошел курс "Django Framework" и я в полном восторге! Погружение в мир веб-разработки было захватывающим. Преподаватели прекрасно объясняли сложные концепции и помогли освоить фреймворки. Теперь я уверенно создаю современные веб-приложения. Большое спасибо за отличное обучение!"',

    },
    {
        name:'Aruzhan Murzabek',
        image:'assets/images/users/3.jpg',
        commentTime:'2023-05-27',
        rating:2,
        comment:'"Курс "JavaScript с нуля" - отличное начало для новичков в программировании! Я получил крепкую базу знаний о синтаксисе, структурах данных и алгоритмах. Материал был представлен понятно и доступно. Теперь я чувствую уверенность в создании своих первых программ. Спасибо за отличное обучение!"',
    
    },
    {
        name:'Daurzhan Tolepbergen',
        rating:5,
        image:'assets/images/users/4.jpg',
        commentTime:'2023-05-27',
        comment:'"Курс "Python для начинающих" - идеальное решение для тех, кто хочет погрузиться в анализ данных. Я изучил мощные инструменты Python, такие как NumPy и Pandas, а также научился визуализировать данные с помощью библиотеки Matplotlib. Этот курс дал мне навыки, необходимые для работы в сфере аналитики данных. Очень рекомендую!"       ',
    

    }

]