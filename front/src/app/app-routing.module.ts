import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { JobListComponent } from './job-list/job-list.component';
import { LoginGuard } from './login/login-guard';
import { LoginComponent } from './login/login.component';

const routes: Routes = [
  { path: '',  pathMatch: 'full', redirectTo: 'doogeun.co/position-list' },
  { path: 'login', component: LoginComponent },
  { path: ':companyName', redirectTo: ':companyName/position-list' },
  { path: ':companyName/position-list', component: JobListComponent },
  { path: ':companyName/position/:positionId', component: JobListComponent },
  {
    path: ':companyName/admin',
    canActivate: [ LoginGuard ],
    loadChildren: () => import('./admin/admin.module').then(m => m.AdminModule),
  },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
