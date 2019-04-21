import {NgModule} from "@angular/core";
import {Routes, RouterModule} from "@angular/router";
import {Error404PageComponent} from "./pages/404-page/error404-page.component";
import {IndexPageComponent} from "./pages/index-page/index-page.component";
import {LoginPageComponent} from "./pages/login-page/login-page.component";

const routes: Routes = [
    {path: "", pathMatch: "full", component: IndexPageComponent},
    {path: "login", component: LoginPageComponent},
    {path: "404", component: Error404PageComponent},
    {path: "**", redirectTo: "/404"}
];

@NgModule({
    imports: [RouterModule.forRoot(routes)],
    exports: [RouterModule]
})
export class AppRoutingModule {
}
