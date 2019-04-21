import {BrowserModule} from "@angular/platform-browser";
import {NgModule} from "@angular/core";

import {AppRoutingModule} from "./app-routing.module";
import {AppComponent} from "./app.component";
import { IndexPageComponent } from './pages/index-page/index-page.component';
import { LoginPageComponent } from './pages/login-page/login-page.component';
import { Error404PageComponent } from './pages/404-page/error404-page.component';

@NgModule({
    declarations: [
        AppComponent,
        IndexPageComponent,
        LoginPageComponent,
        Error404PageComponent
    ],
    imports: [
        BrowserModule,
        AppRoutingModule
    ],
    providers: [],
    bootstrap: [AppComponent]
})
export class AppModule {
}
