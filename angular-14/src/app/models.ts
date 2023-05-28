
export interface Course{
    id:string,
    name:string,
    description:string,
    imgURL:string,
}
export interface User{
    id:string,
    name:string,
    email:string
}
export interface Tokena{
    accessToken:string
}
export interface Order{
    id:string,
    user_id:string,
    course_id:string
}
export interface OrderResponse{
    status:any,
    error:any
}
export interface Review{
    user_name:string
    course_name:string
    rating :number
    comment:string
    created_at:string
}