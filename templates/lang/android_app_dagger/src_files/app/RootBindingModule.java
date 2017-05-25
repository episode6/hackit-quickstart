package {{ .Group }}.app;

import {{ .Group }}.main.MainActivity;
import {{ .Group }}.main.MainActivityModule;

import dagger.Module;
import dagger.android.ContributesAndroidInjector;

@Module
abstract class RootBindingModule {

  @ContributesAndroidInjector(modules = {MainActivityModule.class})
  abstract MainActivity mainActivity();

}
