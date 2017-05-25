package {{ .Group }}.main;

import dagger.Module;
import dagger.android.ContributesAndroidInjector;

@Module
public abstract class MainActivityModule {

  @ContributesAndroidInjector
  abstract MainFragment mainFragment();
}
