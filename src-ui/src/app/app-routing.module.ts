import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { NewsComponent } from './news/news.component';
import { NavBarComponent } from './nav-bar/nav-bar.component';
import { ShortcutBarComponent } from './shortcut-bar/shortcut-bar.component';

const routes: Routes = [
  {
    path: '',
    pathMatch: 'full',
    component: ShortcutBarComponent,
    outlet: 'left'
  },
  {
    path: '',
    pathMatch: 'full',
    component: NewsComponent,
    outlet: 'center'
  },
  {
    path: '',
    pathMatch: 'full',
    component: NavBarComponent,
    outlet: 'right'
  }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
