package {{ .Group }}.main;

import android.app.Activity;

import {{ .Group }}.base.BaseActivityModule;

import com.episode6.hackit.android.inject.context.scope.ActivityScope;

import dagger.Binds;
import dagger.Module;
import dagger.android.ContributesAndroidInjector;

@Module(includes = {BaseActivityModule.class})
public abstract class MainActivityModule {

  @Binds
  @ActivityScope
  abstract Activity bindMainActivity(MainActivity activity);

  @ContributesAndroidInjector
  abstract MainFragment mainFragment();
}
