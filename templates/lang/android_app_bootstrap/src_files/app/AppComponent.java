package {{ .Group }}.app;

import javax.inject.Singleton;

import dagger.Component;
import dagger.android.AndroidInjectionModule;
import dagger.android.AndroidInjector;

@Singleton
@Component(modules = {
    AndroidInjectionModule.class,
    {{ .CamelNameWithoutApp }}AppModule.class,
    RootBindingModule.class})
interface {{ .CamelNameWithoutApp }}AppComponent extends AndroidInjector<{{ .CamelNameWithoutApp }}App> {

  @Component.Builder
  abstract class Builder extends AndroidInjector.Builder<{{ .CamelNameWithoutApp }}App> {}
}
