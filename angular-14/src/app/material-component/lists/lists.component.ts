import { Component } from '@angular/core';
import {MatListModule} from '@angular/material/list';
import { BackendService } from 'src/app/backend.service';
import { Course ,Order} from 'src/app/models';

@Component({
  selector: 'app-lists',
  templateUrl: './lists.component.html',
  styleUrls: ['./lists.component.scss']
})
export class ListsComponent {
  courses:Order[]=[];
  constructor(private back:BackendService){}
  ngOnInit(){
    const token  = localStorage.getItem('token')
    console.log(token)
    if (token){
      this.back.getUserCourses(token).subscribe((val)=>{
        this.courses=val
        console.log(val)
      })
      
    }
    this.getCourses()

  }
  getCourses(){
    const courses= this.back.getUserCourses("asd")
    console.log(courses)
  }
  folders = [
    {
      name: 'Python 3',
      updated: new Date('1/28/04')
    },
    {
      name: 'Django',
      updated: new Date('1/17/22')
    },
    {
      name: 'Angular 15.3',
      updated: new Date('1/28/23')
    }
  ];

}
