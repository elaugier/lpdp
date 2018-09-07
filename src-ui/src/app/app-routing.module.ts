import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { NewsComponent } from './news/news.component';
import { NavBarComponent } from './nav-bar/nav-bar.component';
import { ShortcutBarComponent } from './shortcut-bar/shortcut-bar.component';
import { MyAccountComponent } from './my-account/my-account.component';
import { MySettingsComponent } from './my-settings/my-settings.component';
import { MySecurityComponent } from './my-security/my-security.component';
import { ContentComponent } from './content/content.component';

const routes: Routes = [
  {path: '', redirectTo: 'home', pathMatch: 'full'},
  {path: 'home', component: ContentComponent, children: [
    {path: '', redirectTo: 'news', pathMatch: 'full'},
    {path: 'news', component: NewsComponent},
    {path: 'myaccount', component: MyAccountComponent}
  ]},
  {path: '**', component: ContentComponent}
];

@NgModule({
  imports: [RouterModule.forRoot(routes, { useHash: false, enableTracing: true })],
  exports: [RouterModule]
})
export class AppRoutingModule { }
