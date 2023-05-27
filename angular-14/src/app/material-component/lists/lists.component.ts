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
  courses:Course[]=[];
  constructor(private back:BackendService){}
  ngOnInit(){
    const token  = localStorage.getItem('token')
    console.log(token)
    if (token){
      this.back.getUserCourses(token).subscribe((val)=>{
        this.courses=val
      })

    }
  }
}
