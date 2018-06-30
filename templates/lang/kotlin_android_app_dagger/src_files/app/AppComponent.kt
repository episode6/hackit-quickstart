package {{ .Group }}.app

import dagger.Component
import dagger.android.AndroidInjector
import dagger.android.support.AndroidSupportInjectionModule
import javax.inject.Singleton

@Singleton
@Component(modules = arrayOf(
    AndroidSupportInjectionModule::class,
    {{ .CamelNameWithoutApp }}AppModule::class,
    RootBindingModule::class))
interface {{ .CamelNameWithoutApp }}AppComponent : AndroidInjector<{{ .CamelNameWithoutApp }}App> {

  @Component.Builder
  abstract class Builder : AndroidInjector.Builder<{{ .CamelNameWithoutApp }}App>()
}
