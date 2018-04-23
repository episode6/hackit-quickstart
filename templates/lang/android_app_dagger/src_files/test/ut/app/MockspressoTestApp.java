package {{ .Group }}.app;

import android.app.Activity;
import android.app.Fragment;
import android.app.Service;
import android.content.BroadcastReceiver;
import android.content.ContentProvider;

import com.episode6.hackit.mockspresso.Mockspresso;

import {{ .Group }}.BuildMockspresso;

import org.junit.Before;
import org.mockito.Mock;
import org.mockito.invocation.InvocationOnMock;
import org.mockito.stubbing.Answer;
import org.robolectric.RuntimeEnvironment;

import dagger.android.AndroidInjector;
import dagger.android.DaggerApplication;
import dagger.android.DispatchingAndroidInjector;

import static org.mockito.ArgumentMatchers.any;
import static org.mockito.Mockito.doAnswer;

/**
 * An implementation of {@link {{ .CamelNameWithoutApp }}App} that injects sub-components using
 * Mockspresso. Designed for use with Robolectric.
 */
public class MockspressoTestApp extends {{ .CamelNameWithoutApp }}App {

  public static MockspressoTestApp get() {
    return (MockspressoTestApp) RuntimeEnvironment.application;
  }

  private final DispatchingInjectors mInjectors = new DispatchingInjectors();
  private final Mockspresso mAppMockspresso = BuildMockspresso.withDefaults()
      .testResources(mInjectors)
      .build();

  @Override
  protected AndroidInjector<? extends DaggerApplication> applicationInjector() {
    return new {{ .CamelNameWithoutApp }}AppComponent() {
      @Override
      public void inject({{ .CamelNameWithoutApp }}App instance) {
        mAppMockspresso.inject(instance);
      }
    };
  }

  public Mockspresso.Builder buildUpon() {
    return mAppMockspresso.buildUpon().testResources(new InjectorAttacher());
  }

  class InjectorAttacher {
    @Before
    void setup(Mockspresso mockspresso) {
      mInjectors.setupInjectors(mockspresso);
    }
  }

  private static class DispatchingInjectors {
    @Mock DispatchingAndroidInjector<Activity> activityInjector;
    @Mock DispatchingAndroidInjector<BroadcastReceiver> broadcastReceiverInjector;
    @Mock DispatchingAndroidInjector<Fragment> fragmentInjector;
    @Mock DispatchingAndroidInjector<Service> serviceInjector;
    @Mock DispatchingAndroidInjector<ContentProvider> contentProviderInjector;

    void setupInjectors(final Mockspresso mockspresso) {
      Answer<Object> injectAnswer = new Answer<Object>() {
        @Override
        public Object answer(InvocationOnMock invocation) throws Throwable {
          mockspresso.inject(invocation.getArgument(0));
          return null;
        }
      };
      doAnswer(injectAnswer).when(activityInjector).inject(any(Activity.class));
      doAnswer(injectAnswer).when(broadcastReceiverInjector).inject(any(BroadcastReceiver.class));
      doAnswer(injectAnswer).when(fragmentInjector).inject(any(Fragment.class));
      doAnswer(injectAnswer).when(serviceInjector).inject(any(Service.class));
      doAnswer(injectAnswer).when(contentProviderInjector).inject(any(ContentProvider.class));
    }
  }
}
