package {{ .Group }}.base;

import android.app.Activity;
import android.app.Fragment;
import android.content.Context;
import android.os.Build;

import javax.inject.Inject;

import dagger.android.AndroidInjection;
import dagger.android.AndroidInjector;
import dagger.android.DispatchingAndroidInjector;
import dagger.android.HasFragmentInjector;

public class BaseFragment extends Fragment implements HasFragmentInjector {

  @Inject DispatchingAndroidInjector<Fragment> mChildFragmentInjector;

  @Override
  public void onAttach(Context context) {
    AndroidInjection.inject(this);
    super.onAttach(context);
  }

  @SuppressWarnings("deprecation")
  @Override
  public void onAttach(Activity activity) {
    if (Build.VERSION.SDK_INT < Build.VERSION_CODES.M) {
      AndroidInjection.inject(this);
    }
    super.onAttach(activity);
  }

  @Override
  public AndroidInjector<Fragment> fragmentInjector() {
    return mChildFragmentInjector;
  }
}
