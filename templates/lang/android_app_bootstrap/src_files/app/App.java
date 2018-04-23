package {{ .Group }}.app;

import com.episode6.hackit.chop.Chop;

import javax.inject.Inject;

import dagger.android.AndroidInjector;
import dagger.android.DaggerApplication;

public class {{ .CamelNameWithoutApp }}App extends DaggerApplication {

  @Inject Chop.Tree mChopTree;

  @Override
  public void onCreate() {
    super.onCreate();
    Chop.plantTree(mChopTree);
  }

  @Override
  protected AndroidInjector<? extends DaggerApplication> applicationInjector() {
    return Dagger{{ .CamelNameWithoutApp }}AppComponent.builder().create(this);
  }
}
