package {{ .Group }}.app


import com.episode6.hackit.android.inject.context.scope.ActivityScope
import com.episode6.hackit.android.inject.context.scope.ContextScope
import {{ .Group }}.main.MainActivity
import {{ .Group }}.main.MainActivityModule
import dagger.Module
import dagger.android.ContributesAndroidInjector

@Module
abstract class RootBindingModule {

  @ActivityScope
  @ContextScope
  @ContributesAndroidInjector(modules = arrayOf(MainActivityModule::class))
  internal abstract fun mainActivity(): MainActivity

}
