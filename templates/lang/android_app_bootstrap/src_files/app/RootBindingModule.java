package {{ .Group }}.app;

import {{ .Group }}.main.MainActivity;
import {{ .Group }}.main.MainActivityModule;

import com.episode6.hackit.android.inject.context.scope.ActivityScope;
import com.episode6.hackit.android.inject.context.scope.ContextScope;

import dagger.Module;
import dagger.android.ContributesAndroidInjector;

@Module
abstract class RootBindingModule {

  @ActivityScope
  @ContextScope
  @ContributesAndroidInjector(modules = {MainActivityModule.class})
  abstract MainActivity mainActivity();

}
