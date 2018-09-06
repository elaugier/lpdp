import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { ModuleWithProviders } from '@angular/core';
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
    pathMatch: 'full',
    component: ShortcutBarComponent,
    outlet: 'left'
  },
  {
    path: 'myAccount',
    pathMatch: 'full',
    component: MyAccountComponent,
    outlet: 'center'
  },
  {
    path: 'myAccount',
    pathMatch: 'full',
    component: NavBarComponent,
    outlet: 'right'
  },
  {
    path: 'mySettings',
    pathMatch: 'full',
    component: ShortcutBarComponent,
    outlet: 'left'
  },
  {
    path: 'mySettings',
    pathMatch: 'full',
    component: MySettingsComponent,
    outlet: 'center'
  },
  {
    path: 'mySettings',
    pathMatch: 'full',
    component: NavBarComponent,
    outlet: 'right'
  },
  {
    path: 'mySecurity',
    pathMatch: 'full',
    component: ShortcutBarComponent,
    outlet: 'left'
  },
  {
    path: 'mySecurity',
    pathMatch: 'full',
    component: MySecurityComponent,
    outlet: 'center'
  },
  {
    path: 'mySecurity',
    pathMatch: 'full',
    component: NavBarComponent,
    outlet: 'right'
  }
];

export const routingModule: ModuleWithProviders = RouterModule.forRoot(routes);

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
