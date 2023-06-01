import { Component } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import {MatListModule} from '@angular/material/list';
import { BackendService } from 'src/app/backend.service';
import { ConfirmComponent } from 'src/app/confirm/confirm.component';
import { Course ,Order,CourseWithOrderId} from 'src/app/models';
import { OtpiskaComponent } from 'src/app/otpiska/otpiska.component';
import { Location } from '@angular/common';
import { SuccessComponent } from 'src/app/success/success.component';
import { ReviewDialogComponent } from 'src/app/review-dialog/review-dialog.component';

@Component({
  selector: 'app-lists',
  templateUrl: './lists.component.html',
  styleUrls: ['./lists.component.scss']
})
export class ListsComponent {
  courses:CourseWithOrderId[]=[];
  constructor(private back:BackendService,private dialog: MatDialog,private location: Location){}
  ngOnInit(){
    const token  = localStorage.getItem('token')
    console.log(token)
    if (token){
      this.back.getUserCourses(token).subscribe((val)=>{
        this.courses=val
      })

    }
  }
  testt(orderId:string):void{
    const dialogRef = this.dialog.open(OtpiskaComponent);

    dialogRef.afterClosed().subscribe(result => {
      if (result) {
        const token = localStorage.getItem("token")
        if (token){
          this.back.deleteOrder(token,orderId).subscribe()
          this.refreshPage()
          // this.success()
        }
       
      }
  });
}
refreshPage(): void {
  this.location.go(this.location.path());
  window.location.reload();
}
success(){
  const dialogRef = this.dialog.open(SuccessComponent);

  dialogRef.afterClosed().subscribe();
}

review(orderId:string){
  this.openReviewDialog()
  localStorage.setItem('orderId',orderId)
}
openReviewDialog(): void {
  const dialogRef = this.dialog.open(ReviewDialogComponent, {
    width: '400px', // задайте ширину окна по своему усмотрению
  });
}

}
