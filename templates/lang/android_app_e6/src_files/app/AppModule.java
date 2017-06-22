package {{ .Group }}.app;

import android.app.Application;

import com.episode6.hackit.android.inject.context.module.ApplicationContextModule;

import {{ .Group }}.executor.RootExecutorsModule;
import {{ .Group }}.preference.RootPreferencesModule;

import javax.inject.Singleton;

import dagger.Binds;
import dagger.Module;

@Module(includes = {
    ApplicationContextModule.class,
    DebugAppModule.class,
    RootExecutorsModule.class,
    RootPreferencesModule.class})
abstract class {{ .CamelNameWithoutApp }}AppModule {

  @Binds
  @Singleton
  abstract Application bindApp({{ .CamelNameWithoutApp }}App app);
}
