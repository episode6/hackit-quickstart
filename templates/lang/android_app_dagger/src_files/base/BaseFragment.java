package {{ .Group }}.base;

import android.app.Activity;
import android.content.Context;
import android.os.Build;
import android.support.v4.app.Fragment;

import javax.inject.Inject;

import dagger.android.AndroidInjector;
import dagger.android.DispatchingAndroidInjector;
import dagger.android.support.AndroidSupportInjection;
import dagger.android.support.HasSupportFragmentInjector;

public class BaseFragment extends Fragment implements HasSupportFragmentInjector {

  @Inject DispatchingAndroidInjector<Fragment> mChildFragmentInjector;

  @Override
  public void onAttach(Context context) {
    AndroidSupportInjection.inject(this);
    super.onAttach(context);
  }

  @SuppressWarnings("deprecation")
  @Override
  public void onAttach(Activity activity) {
    if (Build.VERSION.SDK_INT < Build.VERSION_CODES.M) {
      AndroidSupportInjection.inject(this);
    }
    super.onAttach(activity);
  }

  @Override
  public AndroidInjector<Fragment> supportFragmentInjector() {
    return mChildFragmentInjector;
  }
}
