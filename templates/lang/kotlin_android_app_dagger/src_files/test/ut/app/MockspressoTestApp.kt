package {{ .Group }}.app

import {{ .Group }}.BuildMockspresso
import android.app.Activity
import android.app.Service
import android.content.BroadcastReceiver
import android.content.ContentProvider
import android.support.v4.app.Fragment
import com.episode6.hackit.mockspresso.Mockspresso
import com.episode6.hackit.mockspresso.quick.QuickMockspresso
import dagger.android.AndroidInjector
import dagger.android.DispatchingAndroidInjector
import dagger.android.support.DaggerApplication
import org.junit.Before
import org.mockito.ArgumentMatchers.any
import org.mockito.Mock
import org.mockito.Mockito.doAnswer
import org.mockito.stubbing.Answer
import org.robolectric.RuntimeEnvironment

/**
 * An implementation of [{{ .CamelNameWithoutApp }}App] that injects sub-components using
 * Mockspresso. Designed for use with Robolectric.
 */
class MockspressoTestApp : {{ .CamelNameWithoutApp }}App() {

  private val injectors = DispatchingInjectors()
  private val appMockspresso = BuildMockspresso.withDefaults()
      .testResources(injectors)
      .build()

  override fun applicationInjector(): AndroidInjector<out DaggerApplication> {
    return object : {{ .CamelNameWithoutApp }}AppComponent {
      override fun inject(instance: {{ .CamelNameWithoutApp }}App) {
        appMockspresso.inject(instance)
      }
    }
  }

  fun buildUpon(): QuickMockspresso.Builder {
    return appMockspresso.buildUpon().testResources(InjectorAttacher())
  }

  internal inner class InjectorAttacher {
    @Before
    fun setup(mockspresso: Mockspresso) {
      injectors.setupInjectors(mockspresso)
    }
  }

  private class DispatchingInjectors {
    @Mock
    lateinit var activityInjector: DispatchingAndroidInjector<Activity>
    @Mock
    lateinit var broadcastReceiverInjector: DispatchingAndroidInjector<BroadcastReceiver>
    @Mock
    lateinit var fragmentInjector: DispatchingAndroidInjector<Fragment>
    @Mock
    lateinit var serviceInjector: DispatchingAndroidInjector<Service>
    @Mock
    lateinit var contentProviderInjector: DispatchingAndroidInjector<ContentProvider>

    fun setupInjectors(mockspresso: Mockspresso) {
      val injectAnswer = Answer<Any> { invocation ->
        mockspresso.inject(invocation.getArgument(0))
        null
      }

      doAnswer(injectAnswer).`when`(activityInjector).inject(any(Activity::class.java))
      doAnswer(injectAnswer).`when`(broadcastReceiverInjector).inject(any(BroadcastReceiver::class.java))
      doAnswer(injectAnswer).`when`(fragmentInjector).inject(any(Fragment::class.java))
      doAnswer(injectAnswer).`when`(serviceInjector).inject(any(Service::class.java))
      doAnswer(injectAnswer).`when`(contentProviderInjector).inject(any(ContentProvider::class.java))
    }
  }

  companion object {
    fun get(): MockspressoTestApp {
      return RuntimeEnvironment.application as MockspressoTestApp
    }
  }
}
