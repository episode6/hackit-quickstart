package {{ .Group }}.app

import dagger.android.AndroidInjector
import dagger.android.support.DaggerApplication

open class {{ .CamelNameWithoutApp }}App : DaggerApplication() {

  override fun applicationInjector(): AndroidInjector<out DaggerApplication> {
    return Dagger{{ .CamelNameWithoutApp }}AppComponent.builder().create(this);
  }
}
