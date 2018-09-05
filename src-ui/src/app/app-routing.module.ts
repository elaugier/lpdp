import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { NewsComponent } from './news/news.component';
import { NavBarComponent } from './nav-bar/nav-bar.component';
import { ShortcutBarComponent } from './shortcut-bar/shortcut-bar.component';
import { MyAccountComponent } from './my-account/my-account.component';
import { MySettingsComponent } from './my-settings/my-settings.component';
import { MySecurityComponent } from './my-security/my-security.component';

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
  },
  {
    path: 'myAccount',
    component: MyAccountComponent,
    outlet: 'center'
  },
  {
    path: 'mySettings',
    component: MySettingsComponent,
    outlet: 'center'
  },
  {
    path: 'mySecurity',
    component: MySecurityComponent,
    outlet: 'center'
  }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
