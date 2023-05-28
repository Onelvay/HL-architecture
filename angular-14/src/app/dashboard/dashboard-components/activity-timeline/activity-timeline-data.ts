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
        comment:'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Integer nec odio. Praesent libero.',

    },
    {
        name:'Aruzhan Murzabek',
        image:'assets/images/users/3.jpg',
        commentTime:'2023-05-27',
        rating:2,
        comment:'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Integer nec odio. Praesent libero.',
    
    },
    {
        name:'Daurzhan Tolepbergen',
        rating:5,
        image:'assets/images/users/4.jpg',
        commentTime:'2023-05-27',
        comment:'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Integer nec odio. Praesent libero.',
    

    }

]