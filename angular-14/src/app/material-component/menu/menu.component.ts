import { Component } from '@angular/core';
import { MatMenuModule } from '@angular/material/menu';
import { Course } from 'src/app/models';
import { BackendService } from 'src/app/backend.service';
import { MatDialog } from '@angular/material/dialog';
import { ConfirmComponent } from 'src/app/confirm/confirm.component';
import { ErrorComponent } from 'src/app/error/error.component';
@Component({
  selector: 'app-menu',
  templateUrl: './menu.component.html',
  styleUrls: ['./menu.component.scss']
})
export class MenuComponent {
  test:Course[]=[];
  constructor(private service:BackendService,private dialog: MatDialog,private err:MatDialog){}
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
          // alert("вы уже подписаны на этот курс")
          this.error()
        }
      })
    }
  }
  confirm(courseId:string):void{
    const dialogRef = this.dialog.open(ConfirmComponent);

    dialogRef.afterClosed().subscribe(result => {
      if (result) {
        this.register(courseId)
      }
  });
}
error(){
  const dialogRef = this.err.open(ErrorComponent);

  dialogRef.afterClosed().subscribe();
}
}
