import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { FormsModule } from '@angular/forms';
import { MaterialModule } from '@blox/material';

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
import { ContentComponent } from './content/content.component';
import { MyPostsComponent } from './my-posts/my-posts.component';
import { MyCommentsComponent } from './my-comments/my-comments.component';
import { MyWarningsComponent } from './my-warnings/my-warnings.component';
import { MgmtTagsComponent } from './mgmt-tags/mgmt-tags.component';
import { MgmtAccountsComponent } from './mgmt-accounts/mgmt-accounts.component';
import { MgmtCorrectionRequestsComponent } from './mgmt-correction-requests/mgmt-correction-requests.component';
import { MgmtWarningsComponent } from './mgmt-warnings/mgmt-warnings.component';
import { MgmtWarningTemplatesComponent } from './mgmt-warning-templates/mgmt-warning-templates.component';
import { MgmtActivitiesComponent } from './mgmt-activities/mgmt-activities.component';
import { MgmtAlertsComponent } from './mgmt-alerts/mgmt-alerts.component';
import { MgmtBagdesComponent } from './mgmt-bagdes/mgmt-bagdes.component';
import { MgmtBagdeTypesComponent } from './mgmt-bagde-types/mgmt-bagde-types.component';
import { MgmtBadIPAdressesComponent } from './mgmt-bad-ipadresses/mgmt-bad-ipadresses.component';
import { MgmtBooksComponent } from './mgmt-books/mgmt-books.component';
import { MgmtContestsComponent } from './mgmt-contests/mgmt-contests.component';
import { MgmtLicensesComponent } from './mgmt-licenses/mgmt-licenses.component';
import { MgmtPaymentsComponent } from './mgmt-payments/mgmt-payments.component';
import { MgmtVotingsComponent } from './mgmt-votings/mgmt-votings.component';
import { MyPaymentsComponent } from './my-payments/my-payments.component';
import { LoginComponent } from './login/login.component';
import { GraphQLModule } from './graphql.module';
import { HttpClientModule } from '@angular/common/http';

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
    MyCorrectionRequestsComponent,
    ContentComponent,
    MyPostsComponent,
    MyCommentsComponent,
    MyWarningsComponent,
    MgmtTagsComponent,
    MgmtAccountsComponent,
    MgmtCorrectionRequestsComponent,
    MgmtWarningsComponent,
    MgmtWarningTemplatesComponent,
    MgmtActivitiesComponent,
    MgmtAlertsComponent,
    MgmtBagdesComponent,
    MgmtBagdeTypesComponent,
    MgmtBadIPAdressesComponent,
    MgmtBooksComponent,
    MgmtContestsComponent,
    MgmtLicensesComponent,
    MgmtPaymentsComponent,
    MgmtVotingsComponent,
    MyPaymentsComponent,
    LoginComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    FormsModule,      // using FormsModule is optional
    MaterialModule, GraphQLModule, HttpClientModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
