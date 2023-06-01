import { Injectable } from '@angular/core';
import {HttpClient, HttpHeaders} from "@angular/common/http";
import {Observable} from "rxjs";
import { Course,User,Tokena,Order,OrderResponse,Review,CourseWithOrderId ,PostReview} from './models';
import { SassString } from 'sass';

@Injectable({
  providedIn: 'root'
})
export class BackendService {
  headers = { 'Content-Type': 'application/json' };
  BASE_URL = 'http://localhost:8080'
  
  constructor(private client: HttpClient) {}
  getCourses(): Observable<Course[]> {
    return this.client.get<Course[]>(`${this.BASE_URL}/courses`);
  }

  signIn(email: string,password:string){
    return this.client.post<Tokena>(
      `${this.BASE_URL}/auth/sign-in`,
      {email:email, password:password}
    )
  }
  signUp(name:string,password:string){
    return this.client.post<Tokena>(
      `${this.BASE_URL}/auth/sign-up`,
      {name:"default",email:name, password:password}
    )
  }

  getUserCourses(token:string):Observable<CourseWithOrderId[]>{
    const headers = new HttpHeaders().set('Authorization', `Bearer ${token}`);
    return this.client.get<CourseWithOrderId[]>(`${this.BASE_URL}/user/course`,{headers});
  }
  getReviews(token:string):Observable<Review[]>{
    const headers = new HttpHeaders().set('Authorization', `Bearer ${token}`);
    return this.client.get<Review[]>(`${this.BASE_URL}/order/review`,{headers});
  }
  addCourse(courseId:string,token:string):Observable<OrderResponse>{
    const headers = new HttpHeaders().set('Authorization', `Bearer ${token}`);
    return this.client.post<OrderResponse>(
      `${this.BASE_URL}/order`,
      {course_id:courseId},{headers})
  }
  addReview(token:string,order_id:string,comment:string,rating:number){
    const headers = new HttpHeaders().set('Authorization', `Bearer ${token}`);
    
    return this.client.post<OrderResponse>(
      `${this.BASE_URL}/order/review`,
      {order_id:order_id,rating:rating,comment:comment},{headers})
  }

  deleteOrder(token:string,order_id:string){
    const headers = new HttpHeaders().set('Authorization', `Bearer ${token}`);
    return this.client.delete<OrderResponse>(`${this.BASE_URL}/order/${order_id}`,{headers})
  }
}
