package tech.evove.goandroid;

import java.util.concurrent.TimeUnit;

import io.reactivex.Scheduler;
import io.reactivex.annotations.NonNull;
import io.reactivex.disposables.Disposable;
import io.reactivex.internal.disposables.EmptyDisposable;

class GoWorker extends Scheduler.Worker {
    private final core.Worker actual;

    GoWorker(core.Worker actual) {
        this.actual = actual;
    }

    @Override
    public Disposable schedule(@NonNull Runnable runnable, long l, @NonNull TimeUnit timeUnit) {
        if (actual.isDisposed()) {
            return EmptyDisposable.INSTANCE;
        }
        return GoScheduler.schedule(actual, runnable, l, timeUnit);
    }

    @Override
    public void dispose() {
        actual.dispose();
    }

    @Override
    public boolean isDisposed() {
        return actual.isDisposed();
    }
}
