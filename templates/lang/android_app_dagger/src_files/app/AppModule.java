package {{ .Group }}.app;

import android.app.Application;

import dagger.Binds;
import dagger.Module;

@Module
abstract class {{ .CamelNameWithoutApp }}AppModule {

  @Binds
  abstract Application bindApp({{ .CamelNameWithoutApp }}App app);
}
