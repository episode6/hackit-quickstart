package {{ .Group }}.app;

import dagger.android.AndroidInjector;
import dagger.android.support.DaggerApplication;

public class {{ .CamelNameWithoutApp }}App extends DaggerApplication {

  @Override
  protected AndroidInjector<? extends DaggerApplication> applicationInjector() {
    return Dagger{{ .CamelNameWithoutApp }}AppComponent.builder().create(this);
  }
}
