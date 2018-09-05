import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { TopBarComponent } from './top-bar/top-bar.component';
import { MenuBarComponent } from './menu-bar/menu-bar.component';
import { ShowSectionComponent } from './show-section/show-section.component';
import { NewsComponent } from './news/news.component';
import { NavBarComponent } from './nav-bar/nav-bar.component';
import { ShortcutBarComponent } from './shortcut-bar/shortcut-bar.component';
import { MySecurityComponent } from './my-security/my-security.component';
import { MyAccountComponent } from './my-account/my-account.component';
import { MySettingsComponent } from './my-settings/my-settings.component';
import { MyMessagesComponent } from './my-messages/my-messages.component';
import { MyCorrectionRequestsComponent } from './my-correction-requests/my-correction-requests.component';

@NgModule({
  declarations: [
    AppComponent,
    TopBarComponent,
    MenuBarComponent,
    ShowSectionComponent,
    NewsComponent,
    NavBarComponent,
    ShortcutBarComponent,
    MySecurityComponent,
    MyAccountComponent,
    MySettingsComponent,
    MyMessagesComponent,
    MyCorrectionRequestsComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
