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


  questions: Question[] = [
    { text: 'Как зарегистрироваться на курсы?', answer: 'Для регистрации на курсы просто перейдите на страницу регистрации, заполните необходимую информацию и выберите желаемый курс. После успешной регистрации вы получите подтверждение по электронной почте.', showAnswer: false },
    { text: 'Как выбрать подходящий курс?', answer: ' У нас есть разнообразие курсов по программированию, от начинающих до продвинутых. Вы можете просмотреть подробные описания курсов, их продолжительность, требования и отзывы студентов, чтобы выбрать курс, который лучше всего соответствует вашим потребностям и уровню подготовки.', showAnswer: false },
    { text: 'Как получить доступ к материалам курса?', answer: ' После успешной оплаты курса вы получите доступ к нашей онлайн-платформе обучения, где вы сможете получить все необходимые материалы курса, включая лекции, задания, дополнительные материалы и возможность общения с преподавателями и другими студентами.', showAnswer: false }
  ];

  toggleAnswer(question: Question) {
    question.showAnswer = !question.showAnswer;
  }
}
export interface Question {
  text: string;
  answer: string;
  showAnswer:boolean;
}