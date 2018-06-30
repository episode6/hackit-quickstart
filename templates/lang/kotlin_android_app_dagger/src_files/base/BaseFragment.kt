package {{ .Group }}.base

import android.app.Activity;
import android.content.Context;
import android.os.Build;
import android.support.v4.app.Fragment;

import javax.inject.Inject;

import dagger.android.AndroidInjector;
import dagger.android.DispatchingAndroidInjector;
import dagger.android.support.AndroidSupportInjection;
import dagger.android.support.HasSupportFragmentInjector;

import android.app.Activity
import android.content.Context
import android.os.Build
import android.support.v4.app.Fragment
import dagger.android.AndroidInjector
import dagger.android.DispatchingAndroidInjector
import dagger.android.support.AndroidSupportInjection
import dagger.android.support.HasSupportFragmentInjector
import javax.inject.Inject

abstract class BaseFragment : Fragment(), HasSupportFragmentInjector {

  @Inject
  internal lateinit var childFragmentInjector: DispatchingAndroidInjector<Fragment>

  override fun onAttach(context: Context?) {
    AndroidSupportInjection.inject(this)
    super.onAttach(context)
  }

  override fun onAttach(activity: Activity?) {
    if (Build.VERSION.SDK_INT < Build.VERSION_CODES.M) {
      AndroidSupportInjection.inject(this)
    }
    super.onAttach(activity)
  }

  override fun supportFragmentInjector(): AndroidInjector<Fragment> {
    return childFragmentInjector
  }
}
