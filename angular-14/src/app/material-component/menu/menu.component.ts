import { Component } from '@angular/core';
import { MatMenuModule } from '@angular/material/menu';
import { Course } from 'src/app/models';
import { BackendService } from 'src/app/backend.service';
@Component({
  selector: 'app-menu',
  templateUrl: './menu.component.html',
  styleUrls: ['./menu.component.scss']
})
export class MenuComponent {
  test:Course[]=[];
  constructor(private service:BackendService){}
  ngOnInit(){
    this.service.getCourses().subscribe((val)=>{
      this.test=val
    })
  }
}
