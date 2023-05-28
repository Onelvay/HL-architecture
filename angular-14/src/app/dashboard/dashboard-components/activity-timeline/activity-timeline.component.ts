import { Component, OnInit } from '@angular/core';
import { Activity, activities } from './activity-timeline-data';
import { BackendService } from 'src/app/backend.service';
import { Review } from 'src/app/models';

@Component({
  selector: 'app-activity-timeline',
  templateUrl: './activity-timeline.component.html'
})
export class ActivityTimelineComponent implements OnInit {

  activityData: Activity[];
  comments:Review[]=[];
  constructor(private back:BackendService) {
    this.activityData = activities;
  }


  ngOnInit(): void {
    const token = localStorage.getItem('token')
    if (token){
      this.back.getReviews(token).subscribe((data)=>{
        this.comments=data        
      })
    }
    
  }

}
