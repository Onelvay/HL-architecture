import { Component, OnInit } from '@angular/core';
import {ViewChild } from "@angular/core";
import { Course } from '../models';
import { BackendService } from '../backend.service';
import { Activity, activities } from './activity-timeline-data';

import {
  ApexChart,
  ChartComponent,
  ApexDataLabels,
  ApexPlotOptions,
  ApexLegend,
  ApexTooltip,
  ApexNonAxisChartSeries,
  ApexResponsive,
} from "ng-apexcharts";

export interface VisitorChartOptions {
  series: ApexNonAxisChartSeries;
  chart: ApexChart;
  responsive: ApexResponsive[];
  labels: any;
  tooltip: ApexTooltip;
  legend: ApexLegend;
  colors: string[];
  stroke: any;
  dataLabels: ApexDataLabels;
  plotOptions: ApexPlotOptions;
}

@Component({
  selector: 'app-mainpage',
  templateUrl: './mainpage.component.html',
  styleUrls: ['./mainpage.component.css'],
 
}
)


export class MainpageComponent implements OnInit {
  activityData: Activity[];
  test:Course[]=[];
  @ViewChild("visitor-chart") chart2: ChartComponent = Object.create(null);
  public VisitorChartOptions: Partial<VisitorChartOptions>;

  constructor(private service:BackendService){
    this.activityData = activities;
    this.VisitorChartOptions = {
      series: [4, 5, 2, 1],
      chart: {
        type: "donut",
        fontFamily: "Poppins,sans-serif",
        height: 253,
      },
      plotOptions: {
        pie: {
          donut: {
            size: "80px",
          },
        },
      },
      tooltip: {
        fillSeriesColor: false,
      },
      dataLabels: {
        enabled: false,
      },
      stroke: {
        width: 0,
      },
      legend: {
        show: false,
      },
      labels: ["Backend", "Frontend", "DevOps", "Другие"],
      colors: ["#1e88e5", "#26c6da", "#745af2", "#eceff1"],
      responsive: [
        {
          breakpoint: 767,
          options: {
            chart: {
              width: 200,
            },
          },
        },
      ],
    };

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
