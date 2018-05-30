package {{ .Group }}.app

import android.app.Application
import com.episode6.hackit.android.inject.context.module.ApplicationContextModule
import dagger.Binds
import dagger.Module
import javax.inject.Singleton

@Module(includes = arrayOf(ApplicationContextModule::class))
internal abstract class {{ .CamelNameWithoutApp }}AppModule {

  @Binds
  @Singleton
  internal abstract fun bindApp(app: {{ .CamelNameWithoutApp }}App): Application
}
