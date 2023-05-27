import { Component, OnInit } from '@angular/core';
import { Course } from '../models';
import { BackendService } from '../backend.service';
import { Activity, activities } from './activity-timeline-data';

@Component({
  selector: 'app-mainpage',
  templateUrl: './mainpage.component.html',
  styleUrls: ['./mainpage.component.css']
})
export class MainpageComponent implements OnInit {
  activityData: Activity[];
  test:Course[]=[];

  constructor(private service:BackendService){
    this.activityData = activities;

  }
  ngOnInit(){
    this.service.getCourses().subscribe((val)=>{
      this.test=val
    })
  }
  register(courseId:string){
    const token= localStorage.getItem('token')
    console.log(token)
    if(token){
      this.service.addCourse(courseId,token).subscribe((data)=>{
        if (data.error!=null){
          console.log(data.error)
          alert("вы уже подписаны на этот курс")
        }
      })
    }
  }

}
