package {{ .Group }}.main

import android.app.Activity
import com.episode6.hackit.android.inject.context.scope.ActivityScope
import com.episode6.hackit.android.inject.context.scope.FragmentScope
import {{ .Group }}.base.BaseActivityModule
import dagger.Binds
import dagger.Module
import dagger.android.ContributesAndroidInjector

@Module(includes = arrayOf(BaseActivityModule::class))
abstract class MainActivityModule {

  @Binds
  @ActivityScope
  internal abstract fun bindMainActivity(activity: MainActivity): Activity

  @FragmentScope
  @ContributesAndroidInjector
  internal abstract fun mainFragment(): MainFragment
}
