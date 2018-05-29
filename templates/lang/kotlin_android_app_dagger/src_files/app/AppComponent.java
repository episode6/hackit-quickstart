package {{ .Group }}.app;

import javax.inject.Singleton;

import dagger.Component;
import dagger.android.AndroidInjector;
import dagger.android.support.AndroidSupportInjectionModule;

@Singleton
@Component(modules = {
    AndroidSupportInjectionModule.class,
    {{ .CamelNameWithoutApp }}AppModule.class,
    RootBindingModule.class})
interface {{ .CamelNameWithoutApp }}AppComponent extends AndroidInjector<{{ .CamelNameWithoutApp }}App> {

  @Component.Builder
  abstract class Builder extends AndroidInjector.Builder<{{ .CamelNameWithoutApp }}App> {}
}
