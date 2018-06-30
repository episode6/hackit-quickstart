package {{ .Group }}.base

import android.app.Activity
import android.content.Context
import com.episode6.hackit.android.inject.context.module.ScopedContextModule
import com.episode6.hackit.android.inject.context.scope.ContextScope
import dagger.Binds
import dagger.Module

@Module(includes = arrayOf(ScopedContextModule::class))
abstract class BaseActivityModule {

  @Binds
  @ContextScope
  internal abstract fun bindContext(activity: Activity): Context
}
