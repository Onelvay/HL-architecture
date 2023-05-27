import { Routes } from '@angular/router';

import { FullComponent } from './layouts/full/full.component';
import { LoginComponent } from './login/login.component';
import { ProfileComponent } from './profile/profile.component';
import { SignupComponent } from './signup/signup.component';
import { MainpageComponent } from './mainpage/mainpage.component';

export const AppRoutes: Routes = [
  { path:'login',component:LoginComponent},
  { path:'sign-up',component:SignupComponent},
  {
    path: '',
    component: FullComponent,
    
    children: [
      {path:'mainpage',component:MainpageComponent},
      { path:'settings',component:ProfileComponent},
      
      {
        path: '',
        redirectTo: '/dashboard',
        pathMatch: 'full'
      },
      {
        path: '',
        loadChildren:
          () => import('./material-component/material.module').then(m => m.MaterialComponentsModule)
      },
      {
        path: 'dashboard',
        loadChildren: () => import('./dashboard/dashboard.module').then(m => m.DashboardModule)
      }
    ]
  }
];
